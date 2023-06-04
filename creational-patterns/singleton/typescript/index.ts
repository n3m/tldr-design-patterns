class Singleton {
  private static instance: Singleton;

  private constructor() {
    // Private constructor to prevent instantiation from outside
  }

  /**
   * The getInstance method is responsible for creating the instance if it doesn't exist, and returning the existing instance if it does.
   */
  public static getInstance(): Singleton {
    if (!Singleton.instance) {
      Singleton.instance = new Singleton();
    }
    return Singleton.instance;
  }

  public someMethod(): void {
    // Add your methods here
  }
}

// Usage
const instance1 = Singleton.getInstance();
const instance2 = Singleton.getInstance();

console.log(instance1 === instance2); // Output: true