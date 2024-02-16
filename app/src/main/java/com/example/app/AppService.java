package com.example.app;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.CachePut;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;

@Service
public class AppService {
  @Autowired
  private AppRepository repository;

  @SuppressWarnings("null")
  @Cacheable(value = "peanuts", key = "#id")
  public Schemas getPeanutsById(Long id) {
    return repository.findById(id).orElse(null);
  }

  @SuppressWarnings("null")
  @CachePut(value = "peanuts", key = "#peanuts.id")
  public Schemas savePeanuts(Schemas peanuts) {
    return repository.save(peanuts);
  }
}
