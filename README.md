**What i learn**

need to use defer to end the connection of mongo

controllers/: Contains the application's controller logic, managing HTTP requests and responses.
models/: Defines the data structures and interacts with the database.

always open the package using when building a application, eg.mongo-driver

Why Use context.Context?

  Problem: Long-running operations (e.g., database queries) can hang indefinitely if the database is slow or unresponsive.
  Solution: context.Context allows you to set deadlines or timeouts for operations. If the deadline expires, the operation is canceled automatically.
  
  Problem: If an HTTP request is canceled by a client, associated operations (e.g., database queries or downstream API calls) might continue running, wasting resources.
  Solution: With context.Context, you can propagate a cancellation signal to all parts of an operation, ensuring they stop promptly when no longer needed.
  
  Problem: Sharing data between functions often requires passing multiple parameters, leading to cluttered code.
  Solution: context.Context allows you to attach and retrieve key-value pairs that are specific to a request or operation.
  
  Problem: Goroutines might keep running even after their parent operation is completed or canceled.
  Solution: context.Context helps manage goroutines, ensuring they exit when their context is canceled.

Types of Contexts --> have to read more on it
  context.Background: Used as a base context for top-level operations.
  
  context.WithCancel: Returns a context that can be explicitly canceled.
  
  context.WithTimeout: Sets a timeout, after which the context is automatically canceled.
  
  context.WithValue: Attaches key-value pairs to a context.
