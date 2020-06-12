class Seller {
  final String name;
  final String contactNo;
  final String email;
  final String province;
  final String town;
  final String suburb;

  Seller(this.name, this.contactNo, this.email, this.province, this.town,
      this.suburb);

  Map<String, dynamic> toJson() {
    return {
      "Name": name,
      "ContactNo": contactNo,
      "Email": email,
      "Province": province,
      "Town": town,
      "Suburb": suburb
    };
  }
}
