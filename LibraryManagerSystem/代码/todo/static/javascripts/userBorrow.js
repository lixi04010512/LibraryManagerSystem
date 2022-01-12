function load(){
    $.ajax({
        type:"POST",
        url:"todo/backupdatebook",
        success:function(data){
            $("#add").empty();
          for( var i of data.data){
            $("#add").append(`
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
									<input id="roleNames" required="" placeholder="请输入借阅人姓名" value="" name="roleTotal" type="text">
								</div>
                             `
                   )}
          }
     })
}
load();

$('#borrowBook').click(function(){
    //借阅人是否已经借阅
  $.ajax({
    type:"POST",
    url:"todo/selectborrowuser",
    data:{"user_name":$("#roleNames").val()},
    success:function(data){
      if(data.data ==null){
        $.ajax({
            type:"POST",
            url:"todo/insertborrow",
            data:{"book_name":$("#roleName").val(),"user_name":$("#roleNames").val()},
            success:function(){
                alert("借阅成功！");
                window.location="./First.html"
            }
        })
      }else{
        alert("该用户借阅未归还，无法借阅！");
        window.location="./First.html";
      }
    }
  })
})