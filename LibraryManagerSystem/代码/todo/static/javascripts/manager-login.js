$("#login").click(function(){
    $.ajax({
        type:"POST",
        url:"todo/selectloginmanager",
        data:{"manager_name":$("#login-username").val(),"manager_password":$("#login-password").val()},
        success:function(data){
          if(data.data ==null){
            alert("管理员不存在！");
          }else{
              window.location="./index.html";;
          }
        }
    })
})