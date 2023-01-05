function Register() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/register", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    // var formData = new FormData();
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    // formData.append("username", username);
    // formData.append("password", password);
    var jsonData = {
        "username": username,
        "password": password,
    }
    xhr.send(JSON.stringify(jsonData));
    // var responsetext = xhr.responseText;
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status >= 200 && xhr.status < 300) {
                if (confirm("注册成功，点击确认返回登录页面")) {
                    window.location.href = "/login";
                }
            }else{
                if (xhr.status === 403){
                    alert("该用户已存在");
                }
            }
        }
    };
}