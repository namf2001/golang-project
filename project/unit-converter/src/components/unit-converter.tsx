'use client'

import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Converter } from "@/lib/conver"
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select"

export function UnitConverterComponent() {
  const [typeCovert, settypeCovert] = useState("")
  const [fromUnit, setFromUnit] = useState("")
  const [toUnit, setToUnit] = useState("")
  const [result, setResult] = useState<string | number>("")
  const [activeTab, setActiveTab] = useState("")

  const handleConvert = (value: string, fromUnit: string, toUnit: string, type: string) => {
    var result = Converter(value, fromUnit, toUnit, type)
    setResult(result)
    setActiveTab(type)
  }

  return (
    <Card className="w-full max-w-md mx-auto">
      <CardHeader>
        <CardTitle className="text-2xl font-bold text-center">Unit Converter</CardTitle>
      </CardHeader>
      <CardContent>
        <Tabs defaultValue="length" className="w-full">
          <TabsList className="grid w-full grid-cols-3 mb-4">
            <TabsTrigger value="length">Length</TabsTrigger>
            <TabsTrigger value="weight">Weight</TabsTrigger>
            <TabsTrigger value="temperature">Temperature</TabsTrigger>
          </TabsList>
          <TabsContent value={activeTab ? "" : "length"}>
            <form className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="length">Enter the length to convert</Label>
                <Input
                  id="length"
                  placeholder="Enter length"
                  value={typeCovert}
                  onChange={(e) => settypeCovert(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="fromUnit">Unit to Convert from</Label>
                <Select
                  value={fromUnit}
                  onValueChange={(value) => setFromUnit(value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="e.g., meters, feet" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="meters">Meters</SelectItem>
                    <SelectItem value="kilometers">Kilometers</SelectItem>
                    <SelectItem value="centimeters">Centimeters</SelectItem>
                    <SelectItem value="millimeters">Millimeters</SelectItem>
                    <SelectItem value="feet">Feet</SelectItem>
                    <SelectItem value="inches">Inches</SelectItem>
                    <SelectItem value="yards">Yards</SelectItem>
                    <SelectItem value="miles">Miles</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="space-y-2">
                <Label htmlFor="toUnit">Unit to Convert to</Label>
                <Select
                  value={toUnit}
                  onValueChange={(value) => setToUnit(value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="e.g., meters, feet" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="meters">Meters</SelectItem>
                    <SelectItem value="kilometers">Kilometers</SelectItem>
                    <SelectItem value="centimeters">Centimeters</SelectItem>
                    <SelectItem value="millimeters">Millimeters</SelectItem>
                    <SelectItem value="feet">Feet</SelectItem>
                    <SelectItem value="inches">Inches</SelectItem>
                    <SelectItem value="yards">Yards</SelectItem>
                    <SelectItem value="miles">Miles</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <Button className="w-full" onClick={() => {
                handleConvert(typeCovert, fromUnit, toUnit, "length")
              }}>
                Convert
              </Button>
            </form>
          </TabsContent>
          <TabsContent value={activeTab ? "" : "weight"}>
            <form className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="weight">Enter the weight to convert</Label>
                <Input
                  id="weight"
                  placeholder="Enter weight"
                  value={typeCovert}
                  onChange={(e) => settypeCovert(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="fromUnit">Unit to Convert from</Label>
                <Select
                  value={fromUnit}
                  onValueChange={(value) => setFromUnit(value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="e.g., grams, kilograms" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="grams">Grams</SelectItem>
                    <SelectItem value="kilograms">Kilograms</SelectItem>
                    <SelectItem value="milligrams">Milligrams</SelectItem>
                    <SelectItem value="pounds">Pounds</SelectItem>
                    <SelectItem value="ounces">Ounces</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="space-y-2">
                <Label htmlFor="toUnit">Unit to Convert to</Label>
                <Select
                  value={toUnit}
                  onValueChange={(value) => setToUnit(value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="e.g., grams, kilograms" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="grams">Grams</SelectItem>
                    <SelectItem value="kilograms">Kilograms</SelectItem>
                    <SelectItem value="milligrams">Milligrams</SelectItem>
                    <SelectItem value="pounds">Pounds</SelectItem>
                    <SelectItem value="ounces">Ounces</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <Button className="w-full" onClick={() => {
                handleConvert(typeCovert, fromUnit, toUnit, "weight")
              }}>
                Convert
              </Button>
            </form>
          </TabsContent>
          <TabsContent value={activeTab ? "" : "temperature"}>
            <form className="space-y-4">
              <div className="space-y-2">
                <Label htmlFor="temperature">Enter the temperature to convert</Label>
                <Input
                  id="temperature"
                  placeholder="Enter temperature"
                  value={typeCovert}
                  onChange={(e) => settypeCovert(e.target.value)}
                />
              </div>
              <div className="space-y-2">
                <Label htmlFor="fromUnit">Unit to Convert from</Label>
                <Select
                  value={fromUnit}
                  onValueChange={(value) => setFromUnit(value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="e.g., celsius, fahrenheit" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="celsius">Celsius</SelectItem>
                    <SelectItem value="fahrenheit">Fahrenheit</SelectItem>
                    <SelectItem value="kelvin">Kelvin</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <div className="space-y-2">
                <Label htmlFor="toUnit">Unit to Convert to</Label>
                <Select
                  value={toUnit}
                  onValueChange={(value) => setToUnit(value)}
                >
                  <SelectTrigger>
                    <SelectValue placeholder="e.g., celsius, fahrenheit" />
                  </SelectTrigger>
                  <SelectContent>
                    <SelectItem value="celsius">Celsius</SelectItem>
                    <SelectItem value="fahrenheit">Fahrenheit</SelectItem>
                    <SelectItem value="kelvin">Kelvin</SelectItem>
                  </SelectContent>
                </Select>
              </div>
              <Button className="w-full" onClick={() => {
                handleConvert(typeCovert, fromUnit, toUnit, "temperature")
              }}>
                Convert
              </Button>
            </form>
          </TabsContent>
          <TabsContent value={activeTab}>
            <p className="font-semibold mb-2">Result of your calculation</p>
            <p className="text-2xl font-bold mb-4">{typeCovert}{fromUnit} = {result}{toUnit}</p>
            <Button className="w-full" onClick={() => setActiveTab("")}>
              Reset
            </Button>
          </TabsContent>
        </Tabs>
      </CardContent>
    </Card >
  )
}