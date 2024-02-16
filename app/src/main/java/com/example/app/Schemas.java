package com.example.app;

import javax.persistence.*;
import java.io.Serializable;
import io.swagger.v3.oas.annotations.media.Schema;

@Entity
public class Schemas implements Serializable {
  private static final long serialVersionUID = 1L;

  @Id
  @GeneratedValue(strategy = GenerationType.IDENTITY)
  private Long id;

  @Schema(example = "Nome exemplar")
  private String name;

  @Schema(example = "Descrição exemplar")
  private String description;

  // ID
  public Long getId() {
    return id;
  }

  // Nome
  public String getName() {
    return name;
  }
  public void setName(String name) {
    this.name = name;
  }

  // Descrição
  public String getDescription() {
    return description;
  }
  public void setDescription(String description) {
    this.description = description;
  }
}
