function Converter(value:string, fromUnit:string, toUnit:string, type:string) {
    switch (type) {
        case 'length':
            return lengthConverter(value, fromUnit, toUnit)
        case 'weight':
            return weightConverter(value, fromUnit, toUnit)
        case 'temperature':
            return temperatureConverter(value, fromUnit, toUnit)
        default:
            return 'Invalid type'
    }
}

function lengthConverter(value:string, fromUnit:string, toUnit:string) {
    const units = {
        meters: 1,
        kilometers: 1000,
        centimeters: 0.01,
        millimeters: 0.001,
        feet: 0.3048,
        inches: 0.0254,
        yards: 0.9144,
        miles: 1609.34,
    }
    return convert(value, fromUnit, toUnit, units)
}

function weightConverter(value:string, fromUnit:string, toUnit:string) {
    const units = {
        grams: 1,
        kilograms: 1000,
        milligrams: 0.001,
        pounds: 453.592,
        ounces: 28.3495,
    }
    return convert(value, fromUnit, toUnit, units)
}

function temperatureConverter(value:string, fromUnit:string, toUnit:string) {
    const units = {
        celsius: (value: number) => value,
        fahrenheit: (value: number) => (value - 32) * (5 / 9),
        kelvin: (value: number) => value - 273.15,
    }
    return convert(value, fromUnit, toUnit, units)
}

function convert(value:string, fromUnit:string, toUnit:string, units:any) {
    if (units[fromUnit] && units[toUnit]) {
        return parseFloat(value) * units[fromUnit] / units[toUnit]
    }
    return 'Invalid units'
}

export { Converter }