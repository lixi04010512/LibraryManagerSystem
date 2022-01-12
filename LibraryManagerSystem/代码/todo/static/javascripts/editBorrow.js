//获取编辑对象的信息,显示在页面
function load(){
    $.ajax({
        type:"POST",
        url:"todo/backupdateborrow",
        success:function(data){
            $("#edit").empty();
          for( var i of data.data){
            $("#edit").append(`
        <label>
            借阅id
        </label>
        <div>
            <input id="roleId" required="" placeholder="请输入借阅id" value="${i.borrow_id}" name="roleId" type="text">
        </div> 
        <label>
            图书名称
        </label>
        <div>
            <input id="roleName" required="" placeholder="请输入图书名称" value="${i.book_name}" name="roleName" type="text">
        </div>
        <label>
            借阅人姓名
        </label>
        <div>
            <input id="roleNames" required="" placeholder="请输入借阅人姓名" value="${i.user_name}" name="roleAuthor" type="text">
        </div>
       
              `
            )}
          
        }
    })
}
load();

//修改借阅
$("#editBorrow").click(function(){
    $.ajax({
        type:"POST",
        url:"todo/updateborrow",
        data:{"borrow_id":$("#roleId").val(),"book_name":$("#roleName").val(),"user_name":$("#roleNames").val()},
        success:function(){
           alert("成功修改！");
           window.location="./borrow.html";
        }
    })
})