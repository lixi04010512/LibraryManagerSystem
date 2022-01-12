function load(){
    $.ajax({
        type:"POST",
        url:"/todo/selectbooks",
        success:function(data){
            $("#books_body").empty();
            for( var i of data.data){
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
                            <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnBorrow" data-id="${i.book_id}"><span class="am-icon-pencil-square-o"></span>借阅</button>
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
$("#books_body").delegate(".btnBorrow","click",function(){
    window.location="./userBorrow.html"
    $.ajax({
        type:"POST",
        url:"todo/selectupdatebook",
        data:{"book_id":$(this).data("id")},
        success:function(){
           
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
                    <button type="button" id="role_701" class="am-btn am-btn-default am-btn-xs am-text-secondary btnBorrow" data-id="${i.book_id}"><span class="am-icon-pencil-square-o"></span>借阅</button>
                    </div>
                </div>
            </td>
        </tr>
       `)
     }
   }
  })
})