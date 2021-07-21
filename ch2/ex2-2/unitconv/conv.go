package unitconv

// CToF converte uma temperatura em Celsius para Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converte uma temperatura em Fahrenheit para Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// MeterToFeet converte uma quantidade em metros para pés
func MeterToFeet(m Meter) Feet { return Feet(3.28 * m) }

// FeetToMeter converte uma quantidade em pés para metros
func FeetToMeter(f Feet) Meter { return Meter(0.3 * f) }

// PoundToKilogram converte uma quantidade em libras para kilogramas
func PoundToKilogram(p Pound) Kilogram { return Kilogram(0.45 * p) }

// KilogramToPound converte uma quantidade em kilogramas para libras
func KilogramToPound(k Kilogram) Pound { return Pound(2.2 * k) }
