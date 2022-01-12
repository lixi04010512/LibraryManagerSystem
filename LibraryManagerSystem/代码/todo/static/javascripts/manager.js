//加载所有管理员
function load(){
    $.ajax({
        type:"POST",
        url:"/todo/selectmanagers",
        success:function(data){
            $("#tUser").empty();
            for( var i of data.data){
                $("#tUser").append(`
                <tr>
                <td>${i.manager_id}</td>
                <td>${i.manager_name}</td>
                <td>${i.manager_password}</td>
                <td>
                    <div class="am-btn-toolbar">
                        <div class="am-btn-group am-btn-group-xs">
                            <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.manager_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                            <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.manager_id}"><span class="am-icon-trash-o"></span> 删除</button>
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
        window.location="./editRole.html"
        $.ajax({
            type:"POST",
            url:"todo/selectupdatemanager",
            data:{"manager_id":$(this).data("id")},
            success:function(){
               
            }
        })
})


//删除管理员
$("#tUser").delegate(".btndel","click",function(){
    $.ajax({
        type:"POST",
        url:"todo/deletemanager",
        data:{"manager_id":$(this).data("id")},
        success:function(){
               alert("成功删除！");
               location.reload();
        }
    })
})

//查询单个管理员
$("#btnsearch").click(function(){
    $.ajax({
        type:"POST",
        url:"/todo/selectmanager",
        data:{"manager_name":$("#roleId").val()},
        success:function(data){
            $("#tUser").empty();
            for(var i of data.data){
                $("#tUser").append(`
                <tr>
                <td>${i.manager_id}</td>
                <td>${i.manager_name}</td>
                <td>${i.manager_password}</td>
                <td>
                    <div class="am-btn-toolbar">
                        <div class="am-btn-group am-btn-group-xs">
                            <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.manager_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                            <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.manager_id}"><span class="am-icon-trash-o"></span> 删除</button>
                        </div>
                    </div>
                </td>
            </tr>
           `)
         }
     }
  })
})