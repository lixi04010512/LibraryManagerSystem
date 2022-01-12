//获取编辑对象的信息,显示在页面
function load(){
    $.ajax({
        type:"POST",
        url:"todo/backupdatebook",
        success:function(data){
            $("#edit").empty();
          for( var i of data.data){
            $("#edit").append(`
        <label>
            图书id
        </label>
        <div>
            <input id="roleId" required="" placeholder="请输入图书id" value="${i.book_id}" name="roleId" type="text">
        </div> 
        <label>
            图书名称
        </label>
        <div>
            <input id="roleName" required="" placeholder="请输入图书名称" value="${i.book_name}" name="roleName" type="text">
        </div>
        <label>
            图书作者
        </label>
        <div>
            <input id="roleAuthor" required="" placeholder="请输入图书作者" value="${i.book_author}" name="roleAuthor" type="text">
        </div>
        <label>
            图书价钱
        </label>
        <div>
            <input id="rolePrice" required="" placeholder="请输入图书价钱" value="${i.book_price}" name="rolePrice" type="text">
        </div>
        <label>
            图书位置
        </label>
        <div>
            <input id="rolePosition" required="" placeholder="请输入图书位置" value="${i.book_position}" name="rolePosition" type="text">
        </div>
        <label>
            图书总量
        </label>
        <div>
            <input id="roleTotal" required="" placeholder="请输入图书总量" value="${i.book_total}" name="roleTotal" type="text">
        </div>
              `
            )}
          
        }
    })
}
load();

//修改图书
$("#editBook").click(function(){
    $.ajax({
        type:"POST",
        url:"todo/updatebook",
        data:{"book_id":$("#roleId").val(),"book_name":$("#roleName").val(),"book_author":$("#roleAuthor").val(),"book_price":$("#rolePrice").val(),"book_total":$("#roleTotal").val(),"book_position":$("#rolePosition").val()},
        success:function(){
           alert("成功修改！");
           window.location="./books.html";
        }
    })
})