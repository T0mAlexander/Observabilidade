package com.example.app;

import org.springframework.data.jpa.repository.JpaRepository;

public interface AppRepository extends JpaRepository<AppSchemas, Long> {}