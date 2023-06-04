package singleton

// the singleton struct is unexported (lowercase), making it accessible only within the package. This helps enforce the single instance rule.
type singleton struct{}

var instance *singleton

func GetInstance() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

// Usage
instance1 := singleton.GetInstance()
instance2 := singleton.GetInstance()

fmt.Println(instance1 == instance2) // Output: true