# pub-fwd
Forward port thought nat using public server

```
                        +--------------------------------------------+                            
                        |         Public server (1.2.3.4)            |                            
               +------- | -public -in 0.0.0.0 4455 -out 0.0.0.0 4444 | -------+                  
               |        +--------------------------------------------+        |                  
    +----------+---------+                                         +----------+---------+        
    |         NAT        |                                         |         NAT        |        
    +----------+---------+                                         +----------+---------+        
               |                                                              |                  
+--------------+-------------+                             +------------------+-----------------+
|        C1 (ssh client)     |                             |          c2 (ssh server)           |
| $ ssh root@1.2.3.4 -p 4455 |                             | -in 1.2.3.4:4444 -out localhost:22 |
+----------------------------+                             +------------------------------------+
```
