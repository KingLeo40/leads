import 'photos.dart';
import 'seller.dart';
import 'vehicle.dart';

class Submission {
  final Seller seller;
  final Vehicle vehicle;
  final Photos photos;

  Submission(this.seller, this.vehicle, this.photos);

  Map<String, dynamic> toJson() {
    return {"Seller": seller, "Vehicle": vehicle, "Photos": photos};
  }
}
