package pokeapi

func (c *Client) GetPokemonPerLocation(locationUrl string) ([]RespShallowPokemon, error) {
	resp, err := c.GetSpecificLocation(locationUrl)
	if err != nil {
		return []RespShallowPokemon{}, err
	}
	pokemon := []RespShallowPokemon{}

	for _, encounters := range resp.PokemonEncounters {
		pokemon = append(pokemon, encounters.Pokemon)
	}

	return pokemon, nil
}
