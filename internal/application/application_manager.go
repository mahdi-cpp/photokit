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
	mu             sync.RWMutex
	users          map[string]*account.User
	userManagers   map[string]*UserManager // Maps user IDs to their UserManager
	AccountManager *account.ClientManager
	ctx            context.Context
	cancel         context.CancelFunc
	wg             sync.WaitGroup
}

func NewApplicationManager() (*AppManager, error) {

	ctx, cancel := context.WithCancel(context.Background())
	manager := &AppManager{
		userManagers: make(map[string]*UserManager),
		users:        make(map[string]*account.User),
		ctx:          ctx,
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
//	//	collection []*build_asset.PHCollection[account.User] `json:"collections"`
//	//}
//
//	ac := account.NewAccountManager()
//	users, err := ac.GetAll()
//	if err != nil {
//		return
//	}
//
//	for _, user := range *users {
//		fmt.Printf("User ID: %d\n", user.ID)
//		fmt.Printf("Username: %s\n", user.Username)
//		fmt.Printf("Name: %s %s\n", user.FirstName, user.LastName)
//		manager.users[user.ID] = &user
//	}
//}

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

	var user = manager.users[userID]
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
	// Access users in a thread-safe manner
	users := m.AccountManager.Users // Assuming you add this method to account manager

	fmt.Println("")
	log.Printf("Received user list update with %d users", len(users))

	//Example processing - replace with your actual logic
	//for _, user := range users {
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

	// UpdateOptions application state with the new user list
	m.updateApplicationState(users)
}

func (m *AppManager) updateApplicationState(users []*account.User) {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, user := range users {
		m.users[user.ID] = user
	}

	// UpdateOptions application state based on new user data
	// Example: m.Chats.UpdateUsers(users)
	//m.PrepareAccountChats("0198bdb6-378d-7b3a-8036-a57d3761f5de")

}
