$('#regbtn').click(function(){
    $.ajax({
        type:"POST",
        url:"todo/insertmanager",
        data:{"manager_name":$("#register-username").val(),"manager_password":$("#register-password").val()},
        success:function(data){
            console.log(data);
            alert("注册成功！");
            window.location="./manager-login.html"
        }
    })
})