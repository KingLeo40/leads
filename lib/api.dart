import 'dart:async';
import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/requester.dart';

import 'bodies/submission.dart';

Future<HttpRequest> sendSubmission(Submission obj) async {
  var apiroute = getEndpoint("leads");
  var url = "${apiroute}/submission";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}
