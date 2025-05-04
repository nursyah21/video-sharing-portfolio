package nu.auth.controller;

import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.http.ResponseEntity;

import java.util.HashMap;
import java.util.Map;

import org.springframework.web.bind.annotation.GetMapping;


@RestController
@RequestMapping("/api/v1")
public class Health {
    
    @GetMapping("health")
    public ResponseEntity<Map<String, Object>> checkHealth() {
        Map<String, Object> res = new HashMap<>();
        res.put("status", "ok");

        return ResponseEntity.ok(res);
    }
    
}
