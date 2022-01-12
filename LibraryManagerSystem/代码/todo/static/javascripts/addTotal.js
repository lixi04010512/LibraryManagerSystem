function load(){
    $.ajax({
        type:"POST",
        url:"todo/backaddbook",
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
									补给量
								</label>
								<div>
									<input id="roleTotal" required="" placeholder="请输入补给量" value="" name="roleTotal" type="text">
								</div>
                             `
                   )}
          }
     })
}
load();

$('#addTotal').click(function(){
    $.ajax({
        type:"POST",
        url:"todo/addbook",
        data:{"book_name":$("#roleName").val(),"book_total":$("#roleTotal").val()},
        success:function(){
            alert("补给成功！");
            window.location="./books.html"
        }
    })
})