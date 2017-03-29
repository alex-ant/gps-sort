package flags

import "flag"

// Values contains values of provided execution flags.
var Values struct {
	InputMode string

	InputFile string

	MySQLHost     string
	MySQLPort     int
	MySQLUser     string
	MySQLPass     string
	MySQLDatabase string

	TopAmount int

	ComparisonPointLat float64
	ComparisonPointLng float64
}

func init() {
	// Define flags.
	flag.StringVar(&Values.InputMode, "input-mode", "file", "Data location (file or db)")

	flag.StringVar(&Values.InputFile, "input-file", "geoData.csv", "Input CSV File")

	flag.StringVar(&Values.MySQLHost, "mysql-host", "127.0.0.1", "MySQL host")
	flag.IntVar(&Values.MySQLPort, "mysql-port", 3306, "MySQL port")
	flag.StringVar(&Values.MySQLUser, "mysql-user", "root", "MySQL user")
	flag.StringVar(&Values.MySQLPass, "mysql-pass", "my-secret-pw", "MySQL password")
	flag.StringVar(&Values.MySQLDatabase, "mysql-database", "locations", "MySQL database name")

	flag.IntVar(&Values.TopAmount, "top-amount", 5, "A number of records to show in TOPs")

	flag.Float64Var(&Values.ComparisonPointLat, "comparison-point-lat", 51.925146, "The latitude of the point the distance must be calculated to")
	flag.Float64Var(&Values.ComparisonPointLng, "comparison-point-lng", 4.478617, "The longitude of the point the distance must be calculated to")

	// Parse flags.
	flag.Parse()
}
