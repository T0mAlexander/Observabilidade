package metrics

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Cpu() {
	// Registro global de métricas do Prometheus
	prom := Prometheus()

	// Rotina assíncrona para gerar temperatura fictícia
	go func() {
		// Temporizador de 3 segundos
		timer := time.NewTicker(3 * time.Second)

		for {
			// Gerando algarismo aleatório entre 40 até 75
			randomNumber := rand.Float64()*(75.0-40.0) + 40.0

			// Limitando em 1 casa decimal pós ponto
			generatedNumber := fmt.Sprintf("%.1f", randomNumber)

			// Conversão de string em float64
			temperature, error := strconv.ParseFloat(generatedNumber, 64)

			// Tratamento de erro
			if error != nil {
				fmt.Println("Erro ao converter numeral", error)
			}

			// Informando a métrica fictícia ao Prometheus
			prom.CpuTemperature.Set(temperature)

			// Resetando o temporizador
			<-timer.C
		}
	}()
}
