        type Server map[string]interface{}
        var servers []Server
        err := json.Unmarshal([]byte(out), &servers)package main
