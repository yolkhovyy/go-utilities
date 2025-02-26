# OS Exit Codes

Provided by ChatGPT.

Common generic names for OS exit codes follow standard Unix/Linux conventions. Here are widely recognized exit codes:  

### **🔹 Standard POSIX/Unix Exit Codes**
| Exit Code | Name               | Meaning |
|-----------|--------------------|---------|
| **0**     | `ExitSuccess`      | Successful execution |
| **1**     | `ExitFailure`      | General error / failure |
| **2**     | `ExitUsageError`   | Command-line usage error (invalid args, missing flags) |
| **126**   | `ExitPermission`   | Command found but not executable |
| **127**   | `ExitNotFound`     | Command not found |
| **128**   | `ExitSignalBase`   | Exit due to a signal (e.g., `128 + SIGKILL`) |
| **130**   | `ExitInterrupted`  | Process interrupted (`Ctrl+C`, `SIGINT`) |
| **137**   | `ExitKilled`       | Process killed (`SIGKILL`) |
| **139**   | `ExitSegfault`     | Segmentation fault (`SIGSEGV`) |

---

### **🔹 Additional Common Exit Codes**
| Exit Code | Name                   | Meaning |
|-----------|------------------------|---------|
| **3**     | `ExitConfigError`       | Configuration error |
| **4**     | `ExitDatabaseError`     | Database connection failure |
| **5**     | `ExitIOError`           | Input/output error |
| **6**     | `ExitTimeout`           | Process timeout |
| **7**     | `ExitNetworkError`      | Network-related failure |
| **8**     | `ExitAuthError`         | Authentication failure |

Would you like a more opinionated list for Go-specific applications?
