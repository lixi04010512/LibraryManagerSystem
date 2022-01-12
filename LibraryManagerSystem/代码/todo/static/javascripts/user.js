//加载所有用户
function load(){
$.ajax({
    type:"POST",
    url:"/todo/selectusers",
    success:function(data){
        $("#tUser").empty();
        for(var i of data.data){
            $("#tUser").append(`
            <tr>
            <td>${i.user_id}</td>
            <td>${i.user_name}</td>
            <td>${i.user_password}</td>
            <td>
                <div class="am-btn-toolbar">
                    <div class="am-btn-group am-btn-group-xs">
                        <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.user_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                        <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.user_id}"><span class="am-icon-trash-o"></span> 删除</button>
                    </div>
                </div>
            </td>
        </tr>
            `)
        }
    }
  })
}
load();

//弹出修改页面
$("#tUser").delegate(".btnedit","click",function(){
    window.location="./editUser.html"
    $.ajax({
        type:"POST",
        url:"todo/selectupdateuser",
        data:{"user_id":$(this).data("id")},
        success:function(){
           
        }
    })
})

//删除用户
$("#tUser").delegate(".btndel","click",function(){
    $.ajax({
        type:"POST",
        url:"todo/deleteuser",
        data:{"user_id":$(this).data("id")},
        success:function(){
            alert("成功删除！");
            location.reload();
        }
    })
})

//查询单个用户
$("#btnsearch").click(function(){
    $.ajax({
        type:"POST",
        url:"/todo/selectuser",
        data:{"user_name":$("#roleId").val()},
        success:function(data){
            $("#tUser").empty();
            for(var i of data.data){
                $("#tUser").append(`
            <tr>
            <td>${i.user_id}</td>
            <td>${i.user_name}</td>
            <td>${i.user_password}</td>
            <td>
                <div class="am-btn-toolbar">
                    <div class="am-btn-group am-btn-group-xs">
                        <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.user_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                        <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only" data-id="${i.user_id}"><span class="am-icon-trash-o"></span> 删除</button>
                    </div>
                </div>
            </td>
        </tr>
           `)
         }
     }
  })
})