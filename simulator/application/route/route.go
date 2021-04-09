package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type Route struct {
	ID 			string `json:"routeId"`
	ClientID 	string `json:"clientId"`
	Positions 	[]Position `json:"position"`
}

type Position struct {
	Lat 	float64 `json:"lat"`
	Long 	float64 `json:"long"`
}

type PartialRoutePosition struct {
	ID 			string		`json:"id"`
	ClientID 	string		`json:"clientId"`
	Position 	[]float64	`json:"position"`
	Finished 	bool		`json:"finished"`
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route id not informed")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}

	defer f.Close()

	scanner:= bufio.NewScanner(f)
	for scanner.Scan()	{
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil { return nil }
		long, err := strconv.ParseFloat(data[0], 64)
		if err != nil { return nil }

		r.Positions = append(r.Positions, Position{
			Lat: lat,
			Long: long,
		})
	}

	return nil
}

// Vou pegar um JSON, e colocar numa lista de Strings
func (r *Route) ExportJsonPositions()  ([]string, error){
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	// Vou percorrer todas as minhas posicoes
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		// Se essa for a ultima posicao dps q percorrer positions,
		// vai marcar o route como Finished, pq precisamos passar o finished true ao terminar
		if total-1 == k {
			route.Finished = true
		}
		// Pego uma struct e converto para uma JSON, no caso, route.
		jsonRoute, err := json.Marshal(route)
		// Se triver erro, retorna lista de String vazia e o erro
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}