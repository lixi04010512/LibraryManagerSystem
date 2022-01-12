$("#login").click(function(){
    $.ajax({
        type:"POST",
        url:"todo/selectloginuser",
        data:{"user_name":$("#login-username").val(),"user_password":$("#login-password").val()},
        success:function(data){
          if(data.data ==null){
            alert("用户不存在！");
          }else{
              window.location="./index1.html";
          }
        }
    })
})