package txt

// StatesByCountry maps state names by country code.
var StatesByCountry = LookupTableMap{
	"br": StatesBR,
	"ca": StatesCA,
	"de": StatesDE,
	"us": StatesUS,
}

// StatesBR maps common abbreviations for Brazilian states.
var StatesBR = LookupTable{
	"AC": "Acre",
	"AL": "Alagoas",
	"AM": "Amazonas",
	"AP": "Amapá",
	"BA": "Bahia",
	"CE": "Ceará",
	"DF": "Distrito Federal",
	"ES": "Espírito Santo",
	"GO": "Goiás",
	"MA": "Maranhão",
	"MG": "Minas Gerais",
	"MS": "Mato Grosso do Sul",
	"MT": "Mato Grosso",
	"PA": "Pará",
	"PB": "Paraíba",
	"PE": "Pernambuco",
	"PI": "Piauí",
	"PR": "Paraná",
	"RJ": "Rio de Janeiro",
	"RN": "Rio Grande do Norte",
	"RO": "Rondônia",
	"RR": "Roraima",
	"RS": "Rio Grande do Sul",
	"SC": "Santa Catarina",
	"SE": "Sergipe",
	"SP": "São Paulo",
	"TO": "Tocantins",
}

// StatesCA maps common abbreviations for Canadian provinces and territories.
var StatesCA = LookupTable{
	"AB": "Alberta",
	"BC": "British Columbia",
	"NB": "New Brunswick",
	"NL": "Newfoundland and Labrador",
	"NS": "Nova Scotia",
	"NT": "Northwest Territories",
	"NU": "Nunavut",
	"MB": "Manitoba",
	"ON": "Ontario",
	"PE": "Prince Edward Island",
	"QC": "Quebec",
	"SK": "Saskatchewan",
	"YT": "Yukon",
}

// StatesDE maps common abbreviations for German states.
var StatesDE = LookupTable{
	"BW": "Baden-Württemberg",
	"BY": "Bayern",
	"BE": "Berlin",
	"BB": "Brandenburg",
	"HB": "Bremen",
	"HH": "Hamburg",
	"HE": "Hessen",
	"NI": "Niedersachsen",
	"MV": "Mecklenburg-Vorpommern",
	"NW": "Nordrhein-Westfalen",
	"RP": "Rheinland-Pfalz",
	"SL": "Saarland",
	"SN": "Sachsen",
	"ST": "Sachsen-Anhalt",
	"SH": "Schleswig-Holstein",
	"TH": "Thüringen",
}

// StatesUS maps common abbreviations for US states.
var StatesUS = LookupTable{
	"AL": "Alabama",
	"AK": "Alaska",
	"AS": "American Samoa",
	"AZ": "Arizona",
	"AR": "Arkansas",
	"CA": "California",
	"CO": "Colorado",
	"CT": "Connecticut",
	"DE": "Delaware",
	"DC": "District of Columbia",
	"FM": "Federated States of Micronesia",
	"FL": "Florida",
	"GA": "Georgia",
	"GU": "Guam",
	"HI": "Hawaii",
	"ID": "Idaho",
	"IL": "Illinois",
	"IN": "Indiana",
	"IA": "Iowa",
	"KS": "Kansas",
	"KY": "Kentucky",
	"LA": "Louisiana",
	"MB": "Manitoba",
	"ME": "Maine",
	"MH": "Marshall Islands",
	"MD": "Maryland",
	"MA": "Massachusetts",
	"MI": "Michigan",
	"MN": "Minnesota",
	"MS": "Mississippi",
	"MO": "Missouri",
	"MT": "Montana",
	"NE": "Nebraska",
	"NV": "Nevada",
	"NH": "New Hampshire",
	"NJ": "New Jersey",
	"NM": "New Mexico",
	"NY": "New York",
	"NC": "North Carolina",
	"ND": "North Dakota",
	"NS": "Nova Scotia",
	"MP": "Northern Mariana Islands",
	"OH": "Ohio",
	"OK": "Oklahoma",
	"OR": "Oregon",
	"PW": "Palau",
	"PA": "Pennsylvania",
	"PR": "Puerto Rico",
	"RI": "Rhode Island",
	"SK": "Saskatchewan",
	"SC": "South Carolina",
	"SD": "South Dakota",
	"TN": "Tennessee",
	"TX": "Texas",
	"UT": "Utah",
	"VT": "Vermont",
	"VI": "Virgin Islands",
	"VA": "Virginia",
	"WA": "Washington",
	"WV": "West Virginia",
	"WI": "Wisconsin",
	"WY": "Wyoming",
}
