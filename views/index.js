function Attack() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/index/attack", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    // var formData = new FormData();
    var bossid = document.getElementById("bossid").value;
    var atkval = document.getElementById("atkval").value;
    // formData.append("username", username);
    // formData.append("password", password);
    var jsonData = {
        "bossid": bossid,
        "atkval": atkval,
    }
    xhr.send(JSON.stringify(jsonData));
    // var responsetext = xhr.responseText;
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 201) {
                alert("出刀完成")
                location.reload();
            }
            if (xhr.status === 202) {
                alert("出尾刀完成")
                location.reload();
            }
            if (xhr.status === 415){
                alert("输入数值有误");
            }
        }
    }
}

function Correction() {
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/index/correction", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    // var formData = new FormData();
    var bossid = document.getElementById("bossidc").value;
    var bosssyume = document.getElementById("bosssyume").value;
    var bossvalue = document.getElementById("bossvalue").value;
    // formData.append("username", username);
    // formData.append("password", password);
    var jsonData = {
        "bossid": bossid,
        "bosssyume": bosssyume,
        "bossvalue": bossvalue,
    }
    xhr.send(JSON.stringify(jsonData));
    // var responsetext = xhr.responseText;
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                alert("boss状态调整成功")
                location.reload();
            }

            if (xhr.status === 403){
                alert("用户权限不足");
            }
            if (xhr.status === 415){
                alert("输入数值有误");
            }
        }
    }
}

function LogOut(){
    const xhr = new XMLHttpRequest();
    xhr.open("GET", "/index/logout", true);
    xhr.send();
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status >= 200 && xhr.status < 300 ) {
                alert("退出登录成功")
                location.reload();
            }
        }
    }
}

function ProgressBar(boss_value,now_max_value){
    var b = parseInt(now_max_value)
    var a = parseInt(boss_value)
    return Math.ceil(Number(a) / Number(b) * 100)
}

function GuaShu(bossid){
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/index/guashu", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    var jsonData = {
        "bossid": bossid,
    }
    xhr.send(JSON.stringify(jsonData));
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                alert("挂树成功")
                location.reload();
            }
            if (xhr.status === 400){
                alert("请勿重复挂树");
            }
        }
    }
}

function JinRu(bossid){
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/index/jinru", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    var jsonData = {
        "bossid": bossid,
    }
    xhr.send(JSON.stringify(jsonData));
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                alert("申请攻击boss成功")
                location.reload();
            }
            if (xhr.status === 400){
                alert("有人在打啦");
            }
        }
    }
}

function JinRuRevocation(bossid){
    const xhr = new XMLHttpRequest();
    xhr.open("POST", "/index/jinru/revocation", true);
    xhr.setRequestHeader("Content-Type", "application/json");
    var jsonData = {
        "bossid": bossid,
    }
    xhr.send(JSON.stringify(jsonData));
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            if (xhr.status === 200) {
                alert("撤销成功")
                location.reload();
            }
            if (xhr.status === 400){
                alert("不能撤销他人的申请哦");
            }
        }
    }
}