package application

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/go-account-service/account"
	"github.com/redis/go-redis/v9"
	"log"
	"sync"
	"time"
)

type AppManager struct {
	mu                  sync.RWMutex
	Users               map[string]*account.User
	userManagers        map[string]*UserManager // Maps user IDs to their UserManager
	AccountManager      *account.ClientManager
	ctx                 context.Context
	cancel              context.CancelFunc
	wg                  sync.WaitGroup
	initialUserList     chan struct{} // Add this channel
	initialListReceived bool          // Add this flag
}

func NewApplicationManager() (*AppManager, error) {

	ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	manager := &AppManager{
		userManagers:    make(map[string]*UserManager),
		Users:           make(map[string]*account.User),
		ctx:             ctx,
		cancel:          cancel, // Store the cancel function
		initialUserList: make(chan struct{}),
	}

	var err error

	// Initialize account manager
	manager.AccountManager, err = account.NewClientManager()
	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to initialize account manager: %w", err)
	}

	// Register callback
	manager.AccountManager.Register(manager.accountCallback)

	// Start account operations in a separate goroutine
	manager.wg.Add(1)
	go manager.initAccountOperations()

	return manager, nil
}

//func (manager *AppManager) fetchUsers() {
//
//	//// Alternative using PHCollection directly
//	//type UserPHCollection struct {
//	//	collection []*person_test.PHCollection[account.User] `json:"collections"`
//	//}
//
//	ac := account.NewAccountManager()
//	Users, err := ac.GetAll()
//	if err != nil {
//		return
//	}
//
//	for _, user := range *Users {
//		fmt.Printf("User ID: %d\n", user.ID)
//		fmt.Printf("Username: %s\n", user.Username)
//		fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
//		manager.Users[user.ID] = &user
//	}
//}

func (manager *AppManager) StartUpgrade() error {
	return nil
}

func (manager *AppManager) GetUserManager(c *gin.Context, userID string) (*UserManager, error) {

	manager.mu.Lock()
	defer manager.mu.Unlock()

	// Check if userStorage already exists for this user
	if userManager, exists := manager.userManagers[userID]; exists {
		return userManager, nil
	}

	createManager, err := manager.CreateManager(userID)
	if err != nil {
		return nil, err
	}

	return createManager, nil
}

func (manager *AppManager) CreateManager(userID string) (*UserManager, error) {

	var user = manager.Users[userID]
	if user == nil {
		fmt.Println("user is nil")
		return nil, fmt.Errorf("user not found")
	}

	userManager, err := NewUserManager(user)
	if err != nil {
		return nil, err
	}

	manager.userManagers[userID] = userManager

	return userManager, err
}

func (manager *AppManager) RemoveStorageForUser(userID string) {
	manager.mu.Lock()
	defer manager.mu.Unlock()

	if storage, exists := manager.userManagers[userID]; exists {
		// Cancel any background operations
		storage.cancelMaintenance()
		// Remove from map
		delete(manager.userManagers, userID)
	}
}

func (manager *AppManager) periodicMaintenance() {

	saveTicker := time.NewTicker(10 * time.Second)
	statsTicker := time.NewTicker(30 * time.Minute)
	rebuildTicker := time.NewTicker(24 * time.Hour)
	cleanupTicker := time.NewTicker(1 * time.Hour)

	for {
		select {
		case <-saveTicker.C:
			fmt.Println("saveTicker")
		case <-rebuildTicker.C:
			fmt.Println("rebuildTicker")
		case <-statsTicker.C:
			fmt.Println("statsTicker")
		case <-cleanupTicker.C:
			fmt.Println("cleanupTicker")
		}
	}
}

func (m *AppManager) initAccountOperations() {
	defer m.wg.Done()

	// Request user list with retry logic
	retryCount := 0
	maxRetries := 3
	for {
		if err := m.AccountManager.RequestList(); err != nil {
			log.Printf("Error requesting user list (attempt %d/%d): %v",
				retryCount+1, maxRetries, err)

			retryCount++
			if retryCount >= maxRetries {
				log.Println("Failed to request user list after max retries")
				break
			}

			// Wait before retrying
			select {
			case <-time.After(time.Duration(500*retryCount) * time.Millisecond):
			case <-m.ctx.Done():
				return
			}
			continue
		}
		break
	}

	// Start additional subscribers
	m.AccountManager.StartSubscriber("account/notifications", "account/alerts")
}

func (m *AppManager) accountCallback(msg *redis.Message) {
	switch msg.Channel {
	case "account/list":
		m.processUserList()
	case "account/user":
		// Handle individual user updates
	case "account/notifications":
		// Handle notifications
	case "account/alerts":
		// Handle alerts
	}
}

func (m *AppManager) processUserList() {
	// Access Users in a thread-safe manner
	users := m.AccountManager.Users // Assuming you add this method to account manager

	fmt.Println("")
	log.Printf("Received user list update with %d Users", len(users))

	//Example processing - replace with your actual logic
	//for _, user := range Users {
	//	log.Printf("User: %s %s       (%s)",
	//		user.FirstName, user.LastName, user.PhoneNumber)
	//}

	// Determine the maximum length for First Name + Last Name to ensure consistent spacing
	// You might want to adjust this based on your expected data
	const maxNameLength = 20 // Example: allocate 25 characters for first and last name combined

	for _, user := range users {
		fullName := fmt.Sprintf("   %s %s", user.FirstName, user.LastName)
		log.Printf("%-*s  %s     %s", maxNameLength, fullName, user.PhoneNumber, user.ID)
	}

	for _, user := range users {
		m.Users[user.ID] = user
	}

	// Signal that initial list is received (only once)
	if !m.initialListReceived {
		m.initialListReceived = true
		close(m.initialUserList)
	}

}

func (m *AppManager) WaitForInitialUserList(timeout time.Duration) error {
	select {
	case <-m.initialUserList:
		return nil
	case <-time.After(timeout):
		return fmt.Errorf("timeout waiting for initial user list")
	case <-m.ctx.Done():
		return fmt.Errorf("context cancelled while waiting for user list")
	}
}
