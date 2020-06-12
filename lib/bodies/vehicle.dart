class Vehicle {
  final num year;
  final String make;
  final String makecountry;
  final String model;
  final String trim;
  final String drive;
  final String transmission;
  final String body;
  final String enginepos;
  final num mileage;
  final double price;
  final num condition;
  final String issues;

  Vehicle(
      this.year,
      this.make,
      this.makecountry,
      this.model,
      this.trim,
      this.drive,
      this.transmission,
      this.body,
      this.enginepos,
      this.mileage,
      this.price,
      this.condition,
      this.issues);

  Map<String, dynamic> toJson() {
    return {
      "Year": year,
      "Make": make,
      "MakeCountry": makecountry,
      "Model": model,
      "Trim": trim,
      "Drive": drive,
      "Transmission": transmission,
      "Body": body,
      "EnginePosition": enginepos,
      "Mileage": mileage,
      "Price": price,
      "Condition": condition,
      "Issues": issues
    };
  }
}
