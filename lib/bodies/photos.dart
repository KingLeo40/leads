import 'package:mango_ui/keys.dart';

class Photos {
  final Key front;
  final Key left;
  final Key right;
  final Key rear;
  final Key interior;
  final Key engine;

  Photos(
      this.front, this.left, this.right, this.rear, this.interior, this.engine);

  Map<String, dynamic> toJson() {
    return {
      "Front": front,
      "Left": left,
      "Right": right,
      "Rear": rear,
      "Interior": interior,
      "Engine": engine
    };
  }
}
