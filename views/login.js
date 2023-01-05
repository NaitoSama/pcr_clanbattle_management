function Login() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/login/authentication", true);
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
                alert("登录成功，点击确认进入会战页面")
                window.location.href = "/index";
            }else{
                if (xhr.status === 401){
                    alert("账号或密码错误");
                }
            }
        }
    };
}