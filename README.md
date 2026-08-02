# SolomonNET-Networking project

A distributed wisdom-testing network where the Queen of Sheba challenges Solomon through a secure protocol.  

                    +-------------+    
                    |    Queen    |    
                    +-------------+  
                           |  
                      TCP/TLS  
                           |  
                   +---------------+  
                   | Solomon Server|  
                   +---------------+  
                     /     |      \  
                    /      |       \  
                   /       |        \  
          +---------+ +---------+ +---------+  
          | Witness | | Advisor | | Scribe  |  
          +---------+ +---------+ +---------+

```text
                   Queen (Client)
                        |
                        |
                   TCP Protocol
                        |
                        ▼
                 Solomon Server
                        |
      ┌─────────────────┴─────────────────┐
      │                                   │
 Consensus Engine                    Court Service
      │                                   │
      ▼                                   ▼
 Witnesses                         CaseRepository
      │                                   │
      ▼                                   ▼
 TCP Services                  MemoryRepository
 ```
