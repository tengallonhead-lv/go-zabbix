# Go-Zabbix API Documentation

## Overview

The `go-zabbix` package is a Go library that extends the functionality of the existing Zabbix API client. It provides additional methods for creating hosts, host groups, and items, as well as enhanced item management capabilities.

**Package**: `github.com/tengallonhead-lv/go-zabbix`  
**Go Version**: 1.18+  
**Dependencies**: `github.com/cavaliercoder/go-zabbix`

## Table of Contents

1. [Session Management](#session-management)
2. [Host Management](#host-management)
3. [Host Group Management](#host-group-management)
4. [Item Management](#item-management)
5. [Data Types](#data-types)
6. [Examples](#examples)

---

## Session Management

### ZabbixServer

The `ZabbixServer` struct represents a Zabbix server connection configuration.

```go
type ZabbixServer struct {
    Addr     string `json:"address"`
    User     string `json:"user"`
    Password string `json:"password"`
}
```

**Fields:**
- `Addr`: Zabbix server address (without protocol)
- `User`: Username for authentication
- `Password`: Password for authentication

### Session

The `Session` struct wraps the underlying zabbix.Session and provides extended functionality.

```go
type Session struct {
    *zabbix.Session
}
```

### NewSession() Method

Creates a new Zabbix session with TLS configuration and file-based caching.

```go
func (this *ZabbixServer) NewSession() (*Session, error)
```

**Returns:**
- `*Session`: A new session instance
- `error`: Error if connection fails

**Features:**
- Automatic TLS certificate verification bypass
- File-based session caching (`./zabbix_session`)
- HTTP client with custom transport configuration

---

## Host Management

### CreateHost() Method

Creates a new host in Zabbix.

```go
func (c *Session) CreateHost(params *HostCreateParames) (resp HostCreateResp, err error)
```

**Parameters:**
- `params`: Host creation parameters (see [HostCreateParames](#hostcreateparames))

**Returns:**
- `HostCreateResp`: Response containing created host IDs
- `error`: Error if creation fails

### Supporting Types

#### HostCreateParames

```go
type HostCreateParames struct {
    Host       string        `json:"host"`
    Interfaces []interface{} `json:"interfaces"`
    Groups     []interface{} `json:"groups"`
    Templates  []interface{} `json:"templates,omitempty"`
}
```

**Fields:**
- `Host`: Host name (required)
- `Interfaces`: Array of interface configurations
- `Groups`: Array of host group assignments
- `Templates`: Array of template assignments (optional)

#### HostCreateResp

```go
type HostCreateResp struct {
    HostIds []string `json:"hostids,omitempty"`
}
```

**Fields:**
- `HostIds`: Array of created host IDs

#### Interfaces

```go
type Interfaces struct {
    Type  int    `json:"type"`
    Main  int    `json:"main"`
    Useip int    `json:"useip"`
    IP    string `json:"ip"`
    Dns   string `json:"dns"`
    Port  string `json:"port"`
}
```

**Fields:**
- `Type`: Interface type (1=agent, 2=SNMP, 3=IPMI, 4=JMX)
- `Main`: Main interface flag (0=not main, 1=main)
- `Useip`: Use IP flag (0=use DNS, 1=use IP)
- `IP`: IP address
- `Dns`: DNS name
- `Port`: Port number

#### Groups

```go
type Groups struct {
    Groupid string `json:"groupid"`
}
```

**Fields:**
- `Groupid`: Host group ID

#### Templates

```go
type Templates struct {
    Templateid string `json:"templateid"`
}
```

**Fields:**
- `Templateid`: Template ID

---

## Host Group Management

### CreateHostgroup() Method

Creates a new host group in Zabbix.

```go
func (c *Session) CreateHostgroup(hostgroupName string) (resp HostgroupCreateResponse, err error)
```

**Parameters:**
- `hostgroupName`: Name of the host group to create

**Returns:**
- `HostgroupCreateResponse`: Response containing created group IDs
- `error`: Error if creation fails

### Supporting Types

#### HostgroupCreateParams

```go
type HostgroupCreateParams struct {
    Name string `json:"name"`
}
```

**Fields:**
- `Name`: Host group name

#### HostgroupCreateResponse

```go
type HostgroupCreateResponse struct {
    GroupIds []string `json:"groupids"`
}
```

**Fields:**
- `GroupIds`: Array of created group IDs

---

## Item Management

### GetItems() Method

Retrieves items from Zabbix based on search parameters.

```go
func (c *Session) GetItems(params zabbix.ItemGetParams) ([]Item, error)
```

**Parameters:**
- `params`: Item search parameters from the underlying zabbix library

**Returns:**
- `[]Item`: Array of items matching the search criteria
- `error`: Error if retrieval fails or no items found

### ItemCreate() Method

Creates new monitoring items in Zabbix.

```go
func (c *Session) ItemCreate(params []ItemCreateParams) (resp ItemResponse, err error)
```

**Parameters:**
- `params`: Array of item creation parameters

**Returns:**
- `ItemResponse`: Response containing created item IDs
- `error`: Error if creation fails

### ItemDelete() Method

Deletes monitoring items from Zabbix.

```go
func (c *Session) ItemDelete(itemids []string) (resp ItemResponse, err error)
```

**Parameters:**
- `itemids`: Array of item IDs to delete

**Returns:**
- `ItemResponse`: Response containing deleted item IDs
- `error`: Error if deletion fails

### Supporting Types

#### Item

```go
type Item struct {
    HostID        int    // Unique ID of the Host
    ItemID        int    // Unique ID of the Item
    ItemName      string // Technical name of the Item
    Key_          string // Key of the Item
    ItemDescr     string // Description of the Item
    LastClock     int    // Last Item epoch time
    LastValue     string // Last value of the Item
    LastValueType int    // Type of LastValue (0=float, 1=text, 3=int)
}
```

#### ItemCreateParams

```go
type ItemCreateParams struct {
    zabbix.GetParameters
    
    Name        string   `json:"name,omitempty"`        // Item name
    Key_        string   `json:"key_,omitempty"`        // Item key
    HostID      string   `json:"hostid,omitempty"`      // Host ID
    Type        int      `json:"type,omitempty"`        // Item type
    ValueType   int      `json:"value_type"`            // Value type
    InterfaceID string   `json:"interfaceids,omitempty"` // Interface ID
    Delay       string   `json:"delay,omitempty"`       // Update interval
    
    // Filters
    GraphIDs       []string `json:"graphids,omitempty"`
    TriggerIDs     []string `json:"triggerids,omitempty"`
    ApplicationIDs []string `json:"applicationids,omitempty"`
    
    // Flags
    WebItems     bool   `json:"webitems,omitempty"`
    Inherited    bool   `json:"inherited,omitempty"`
    Templated    bool   `json:"templated,omitempty"`
    Monitored    bool   `json:"monitored,omitempty"`
    WithTriggers bool   `json:"with_triggers,omitempty"`
    
    // String filters
    Group       string `json:"group,omitempty"`
    Host        string `json:"host,omitempty"`
    Application string `json:"application,omitempty"`
}
```

#### ItemResponse

```go
type ItemResponse struct {
    ItemIds []string `json:"itemids,omitempty"`
}
```

**Fields:**
- `ItemIds`: Array of item IDs

---

## Data Types

### Value Types

For `ItemCreateParams.ValueType`:
- `0`: Float
- `1`: Character/String
- `2`: Log
- `3`: Integer
- `4`: Text

### Item Types

For `ItemCreateParams.Type`:
- `0`: Zabbix agent
- `1`: SNMPv1 agent
- `2`: Zabbix trapper
- `3`: Simple check
- `4`: SNMPv2 agent
- `5`: Zabbix internal
- `6`: SNMPv3 agent
- `7`: Zabbix agent (active)
- `8`: Zabbix aggregate
- `9`: Web item
- `10`: External check
- `11`: Database monitor
- `12`: IPMI agent
- `13`: SSH agent
- `14`: TELNET agent
- `15`: Calculated
- `16`: JMX agent
- `17`: SNMP trap

---

## Examples

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    "github.com/tengallonhead-lv/go-zabbix"
)

func main() {
    // Configure Zabbix server
    server := &go_zabbix.ZabbixServer{
        Addr:     "your-zabbix-server.com",
        User:     "admin",
        Password: "zabbix",
    }
    
    // Create session
    session, err := server.NewSession()
    if err != nil {
        log.Fatal("Failed to create session:", err)
    }
    
    // Your API calls here...
}
```

### Creating a Host Group

```go
// Create a host group
response, err := session.CreateHostgroup("My New Host Group")
if err != nil {
    log.Fatal("Failed to create host group:", err)
}

fmt.Printf("Created host group with IDs: %v\n", response.GroupIds)
```

### Creating a Host

```go
// Define host interfaces
interfaces := []interface{}{
    go_zabbix.Interfaces{
        Type:  1,  // Zabbix agent
        Main:  1,  // Main interface
        Useip: 1,  // Use IP
        IP:    "192.168.1.100",
        Port:  "10050",
    },
}

// Define host groups
groups := []interface{}{
    go_zabbix.Groups{
        Groupid: "2", // Linux servers group
    },
}

// Create host parameters
hostParams := &go_zabbix.HostCreateParames{
    Host:       "my-server-01",
    Interfaces: interfaces,
    Groups:     groups,
}

// Create the host
hostResp, err := session.CreateHost(hostParams)
if err != nil {
    log.Fatal("Failed to create host:", err)
}

fmt.Printf("Created host with IDs: %v\n", hostResp.HostIds)
```

### Creating Items

```go
// Create monitoring items
itemParams := []go_zabbix.ItemCreateParams{
    {
        Name:      "CPU Usage",
        Key_:      "system.cpu.util[,idle]",
        HostID:    "10001",
        Type:      0,  // Zabbix agent
        ValueType: 0,  // Float
        Delay:     "30s",
    },
    {
        Name:      "Memory Usage",
        Key_:      "vm.memory.utilization",
        HostID:    "10001",
        Type:      0,  // Zabbix agent
        ValueType: 0,  // Float
        Delay:     "60s",
    },
}

itemResp, err := session.ItemCreate(itemParams)
if err != nil {
    log.Fatal("Failed to create items:", err)
}

fmt.Printf("Created items with IDs: %v\n", itemResp.ItemIds)
```

### Querying Items

```go
import "github.com/cavaliercoder/go-zabbix"

// Query items for a specific host
params := zabbix.ItemGetParams{
    HostIDs: []string{"10001"},
    Output:  "extend",
}

items, err := session.GetItems(params)
if err != nil {
    log.Fatal("Failed to get items:", err)
}

for _, item := range items {
    fmt.Printf("Item: %s (ID: %d, Key: %s, LastValue: %s)\n",
        item.ItemName, item.ItemID, item.Key_, item.LastValue)
}
```

### Deleting Items

```go
// Delete specific items
itemIDs := []string{"10001", "10002", "10003"}

deleteResp, err := session.ItemDelete(itemIDs)
if err != nil {
    log.Fatal("Failed to delete items:", err)
}

fmt.Printf("Deleted items with IDs: %v\n", deleteResp.ItemIds)
```

### Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/tengallonhead-lv/go-zabbix"
    "github.com/cavaliercoder/go-zabbix"
)

func main() {
    // Initialize Zabbix server connection
    server := &go_zabbix.ZabbixServer{
        Addr:     "localhost",
        User:     "Admin",
        Password: "zabbix",
    }
    
    // Create session
    session, err := server.NewSession()
    if err != nil {
        log.Fatal("Failed to create session:", err)
    }
    
    // 1. Create host group
    groupResp, err := session.CreateHostgroup("API Test Group")
    if err != nil {
        log.Fatal("Failed to create host group:", err)
    }
    fmt.Printf("Created host group: %v\n", groupResp.GroupIds)
    
    // 2. Create host
    hostParams := &go_zabbix.HostCreateParames{
        Host: "api-test-host",
        Interfaces: []interface{}{
            go_zabbix.Interfaces{
                Type:  1,
                Main:  1,
                Useip: 1,
                IP:    "192.168.1.100",
                Port:  "10050",
            },
        },
        Groups: []interface{}{
            go_zabbix.Groups{
                Groupid: groupResp.GroupIds[0],
            },
        },
    }
    
    hostResp, err := session.CreateHost(hostParams)
    if err != nil {
        log.Fatal("Failed to create host:", err)
    }
    fmt.Printf("Created host: %v\n", hostResp.HostIds)
    
    // 3. Create items
    itemParams := []go_zabbix.ItemCreateParams{
        {
            Name:      "CPU Usage",
            Key_:      "system.cpu.util[,idle]",
            HostID:    hostResp.HostIds[0],
            Type:      0,
            ValueType: 0,
            Delay:     "30s",
        },
    }
    
    itemResp, err := session.ItemCreate(itemParams)
    if err != nil {
        log.Fatal("Failed to create items:", err)
    }
    fmt.Printf("Created items: %v\n", itemResp.ItemIds)
    
    // 4. Query items
    getParams := zabbix.ItemGetParams{
        HostIDs: []string{hostResp.HostIds[0]},
        Output:  "extend",
    }
    
    items, err := session.GetItems(getParams)
    if err != nil {
        log.Fatal("Failed to get items:", err)
    }
    
    for _, item := range items {
        fmt.Printf("Found item: %s (ID: %d)\n", item.ItemName, item.ItemID)
    }
}
```

---

## Error Handling

The library uses standard Go error handling patterns. All methods return an error as their last return value. Common error scenarios include:

- **Connection failures**: Network issues, invalid credentials
- **API errors**: Invalid parameters, insufficient permissions
- **Resource not found**: When querying non-existent items
- **Validation errors**: Invalid field values or required fields missing

Always check for errors and handle them appropriately in your application.

---

## Best Practices

1. **Session Management**: Reuse sessions when possible to avoid repeated authentication
2. **Error Handling**: Always check and handle errors appropriately
3. **Batch Operations**: Use batch creation for multiple items to improve performance
4. **Resource Cleanup**: Delete unnecessary items and hosts to maintain system performance
5. **Validation**: Validate input parameters before making API calls
6. **Logging**: Use appropriate logging levels for debugging and monitoring

---

## Dependencies

This package depends on:
- `github.com/cavaliercoder/go-zabbix`: Base Zabbix API client
- Go standard library packages: `crypto/tls`, `net/http`, `fmt`, `strconv`

Make sure to install dependencies:
```bash
go mod tidy
```

---

## License

This project extends the open-source Zabbix API libraries and follows the same licensing terms as the base library.