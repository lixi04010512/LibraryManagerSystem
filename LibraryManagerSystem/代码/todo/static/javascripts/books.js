function load(){
    $.ajax({
        type:"POST",
        url:"/todo/selectbooks",
        success:function(data){
            $("#books_body").empty();
            for( var i of data.data){
                console.log(data.data);
                $("#books_body").append(`
                <tr>
                <td>${i.book_id}</td>
                <td>${i.book_name}</td>
                <td>${i.book_author}</td>
                <td class="am-hide-sm-only">${i.book_price}</td>
                <td class="am-hide-sm-only">${i.book_total}</td>
                <td class="am-hide-sm-only">${i.book_position}</td>
                <td>
                    <div class="am-btn-toolbar">
                        <div class="am-btn-group am-btn-group-xs">
                            <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.book_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                            <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.book_id}"><span class="am-icon-trash-o"></span> 删除</button>
                            <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnAdd" data-id="${i.book_id}"><span class="am-icon-pencil-square-o"></span> 补给</button>
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

//新增图书
$(".btnAdd").click(function(){
    window.location="./addBook.html";
})

//弹出修改页面
$("#books_body").delegate(".btnedit","click",function(){
    window.location="./editBook.html"
    $.ajax({
        type:"POST",
        url:"todo/selectupdatebook",
        data:{"book_id":$(this).data("id")},
        success:function(){
           
        }
    })
})

//弹出补给页面
$("#books_body").delegate(".btnAdd","click",function(){
    window.location="./addTotal.html"
    $.ajax({
        type:"POST",
        url:"todo/selectaddbook",
        data:{"book_id":$(this).data("id")},
        success:function(){
           
        }
    })
})


//删除图书
$("#books_body").delegate(".btndel","click",function(){
$.ajax({
    type:"POST",
    url:"todo/deletebook",
    data:{"book_id":$(this).data("id")},
    success:function(){
           alert("成功删除！");
           location.reload();
    }
})
})

//查询单个图书
$(".btnsearch").click(function(){
  $.ajax({
    type:"POST",
    url:"/todo/selectbook",
    data:{"book_name":$("#bookName").val()},
    success:function(data){
        $("#books_body").empty();
        for(var i of data.data){
            $("#books_body").append(`
            <tr>
            <td>${i.book_id}</td>
            <td>${i.book_name}</td>
            <td>${i.book_author}</td>
            <td class="am-hide-sm-only">${i.book_price}</td>
            <td class="am-hide-sm-only">${i.book_total}</td>
            <td class="am-hide-sm-only">${i.book_position}</td>
            <td>
                <div class="am-btn-toolbar">
                    <div class="am-btn-group am-btn-group-xs">
                        <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnedit" data-id="${i.book_id}"><span class="am-icon-pencil-square-o"></span> 编辑</button>
                        <button class="am-btn am-btn-default am-btn-xs am-text-danger am-hide-sm-only btndel" data-id="${i.book_id}"><span class="am-icon-trash-o"></span> 删除</button>
                        <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnAdd" data-id="${i.book_id}"><span class="am-icon-pencil-square-o"></span> 补给</button>
                        </div>
                </div>
            </td>
        </tr>
       `)
     }
 }
})
})