function load(){
    $.ajax({
        type:"POST",
        url:"/todo/selectborrows",
        success:function(data){
            $("#borrow_body").empty();
            for( var i of data.data){
                $("#borrow_body").append(`
                <tr>
                <td>${i.borrow_id}</td>
                <td>${i.book_name}</td>
                <td>${i.user_name}</td>
                <td>
                    <div class="am-btn-toolbar">
                        <div class="am-btn-group am-btn-group-xs">
                        <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.borrow_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                        <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.book_name}"><span class="am-icon-trash-o"></span> 删除</button>
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

//弹出借阅页面
$("#borrow_body").delegate(".btnedit","click",function(){
    window.location="./editBorrow.html"
    $.ajax({
        type:"POST",
        url:"todo/selectupdateborrow",
        data:{"borrow_id":$(this).data("id")},
        success:function(){
           
        }
    })
})

//删除借阅
$("#borrow_body").delegate(".btndel","click",function(){
$.ajax({
    type:"POST",
    url:"todo/deleteborrow",
    data:{"book_name":$(this).data("id")},
    success:function(){
           alert("成功删除！");
           location.reload();
    }
})
})

//查询单个借阅
$(".btnsearch").click(function(){
  $.ajax({
    type:"POST",
    url:"/todo/selectborrow",
    data:{"user_name":$("#username").val()},
    success:function(data){
        $("#borrow_body").empty();
        for(var i of data.data){
            $("#borrow_body").append(`
            <tr>
            <td>${i.borrow_id}</td>
            <td>${i.book_name}</td>
            <td>${i.user_name}</td>
            <td>
                <div class="am-btn-toolbar">
                    <div class="am-btn-group am-btn-group-xs">
                    <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.borrow_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                    <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.borrow_id}"><span class="am-icon-trash-o"></span> 删除</button>
                    </div>
                </div>
            </td>
        </tr>    
       `)
     }
 }
})
})