$(function() {
	$("#login_button").click(function() {
		var username = $("#username").val();
		var password = $("#password").val();
		// Verify the inputted password, so that attackers can't get in
		if(username == "admin" && CryptoJS.SHA256(password).toString() == "cb9a7f035d23c91b81d8f9981405d2e566267e4e10ef8ff0722d3d5a61611b52") {
			// If they have the correct password go to the admin page.
			// For extra security, the hash of the password is added into the name
			window.location.href = "admin_" + CryptoJS.SHA256(password).toString() + ".php";
		}
	})
})