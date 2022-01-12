//获取编辑对象的信息,显示在页面
function load(){
    $.ajax({
        type:"POST",
        url:"todo/backupdatemanager",
        success:function(data){
            $("#edit").empty();
          for( var i of data.data){
            $("#edit").append(`
					<label>
							管理员id
				    </label>
					<div >
						<input id="roleId" required="" placeholder="请输入管理员id" value="${i.manager_id}" name="roleId" type="text">
					</div>
                    <label >
							管理员名称
				    </label>
					<div>
						<input id="roleName" required="" placeholder="请输入管理员名称" value="${i.manager_name}" name="roleName" type="text">
					</div>
                    <label>
                         管理员密码
				    </label>
					<div >
						<input id="rolePassword" required="" placeholder="请输入管理员密码" value="${i.manager_password}" name="rolePassword" type="text">
					</div>
              `
            )}
          
        }
    })
}
load();

//修改管理员
$("#editRole").click(function(){
    $.ajax({
        type:"POST",
        url:"todo/updatemanager",
        data:{"manager_id":$("#roleId").val(),"manager_name":$("#roleName").val(),"manager_password":$("#rolePassword").val()},
        success:function(){
           alert("成功修改！");
           window.location="./manager.html";
        }
    })
})