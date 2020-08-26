import 'package:mango_ui/keys.dart';

import 'seller.dart';

class Submission {
  final Seller seller;
  final Key vehicleKey;

  Submission(this.seller, this.vehicleKey);

  Map<String, dynamic> toJson() {
    return {"Seller": seller, "VehicleKey": vehicleKey};
  }
}
