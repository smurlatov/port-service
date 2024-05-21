package domain

import "fmt"

type Port struct {
	id          string
	name        string
	code        string
	city        string
	country     string
	alias       []string
	regions     []string
	coordinates []float64
	province    string
	timezone    string
	unlocs      []string
}

func NewPort(id, name, code, city, country string, alias, regions []string, coordinates []float64,
	province, timezone string, unlocs []string) (*Port, error) {

	port := &Port{
		id,
		name,
		code,
		city,
		country,
		alias,
		regions,
		coordinates,
		province,
		timezone,
		unlocs,
	}
	//TODO update validation for every field and set all fields by setters
	err := port.SetId(id)
	if err != nil {
		return nil, err
	}

	return port, nil
}

// TODO need to specify validation rulles
func (p *Port) Id() string {
	return p.id
}

func (p *Port) SetId(id string) error {
	if id == "" {
		return fmt.Errorf("%w: port id is required", ErrRequired)
	}
	p.id = id
	return nil
}

func (p *Port) Name() string {
	return p.name
}

func (p *Port) SetName(name string) {
	p.name = name
}

func (p *Port) Code() string {
	return p.code
}

func (p *Port) SetCode(code string) {
	p.code = code
}

func (p *Port) City() string {
	return p.city
}

func (p *Port) SetCity(city string) {
	p.city = city
}

func (p *Port) Country() string {
	return p.country
}

func (p *Port) SetCountry(country string) {
	p.country = country
}

func (p *Port) Alias() []string {
	return p.alias
}

func (p *Port) SetAlias(alias []string) {
	p.alias = alias
}

func (p *Port) Regions() []string {
	return p.regions
}

func (p *Port) SetRegions(regions []string) {
	p.regions = regions
}

func (p *Port) Coordinates() []float64 {
	return p.coordinates
}

func (p *Port) SetCoordinates(coordinates []float64) {
	p.coordinates = coordinates
}

func (p *Port) Province() string {
	return p.province
}

func (p *Port) SetProvince(province string) {
	p.province = province
}

func (p *Port) Timezone() string {
	return p.timezone
}

func (p *Port) SetTimezone(timezone string) {
	p.timezone = timezone
}

func (p *Port) Unlocs() []string {
	return p.unlocs
}

func (p *Port) SetUnlocs(unlocs []string) {
	p.unlocs = unlocs
}
