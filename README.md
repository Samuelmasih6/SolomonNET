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
 # ER Diagram 
```
                  +----------------+
                  |     Cases      |
                  +----------------+
                  | id (PK)        |
                  | question       |
                  | status         |
                  | verdict        |
                  | confidence     |
                  | created_at     |
                  +-------+--------+
                          |
        +------------------+------------------+
        |                  |                  |
        |                  |                  |
+-------v------+   +--------v------+   +-------v--------+
| Testimonies |   |    Events     |   | Advisor Recs   |
+--------------+   +---------------+   +----------------+
| id (PK)      |   | id (PK)       |   | id (PK)        |
| case_id (FK) |   | case_id (FK)  |   | case_id (FK)   |
| witness_idFK |   | event_type    |   | recommendation |
| suspect      |   | payload       |   | confidence     |
| available    |   | created_at    |   | reason         |
+------+-------+   +---------------+   +----------------+
      |
      |
+------v------+
|  Witnesses  |
+-------------+
| id (PK)     |
| name        |
| address     |
| status      |
| reliability |
| last_seen   |
+-------------+
```
