layui.define(['admin'], function (exports) {
    'use strict';
    var admin = layui.admin
        , table = layui.table
        , form = layui.form;

    form.render();

    table.render({
        elem: '#KYO-citys-table'
        , url: '/api/table/citys/paginate'
        , cols: [[
            { field: 'id', width: 100, title: 'ID' }
            , { field: 'province_id', title: '省ID' }
            , { field: 'provinces_province_name', title: '省名称' }
            , { field: 'city_id', title: '市ID' }
            , { field: 'city_name', title: '市名称' }
        ]]
        , page: true
    });

    // 搜索提交
    form.on('submit(KYO-citys-table-search)', function (obj) {
        var field = obj.field; // 获得表单字段
        // 执行搜索重载
        table.reload('KYO-citys-table', {
            page: {
                curr: 1 // 重新从第 1 页开始
            },
            where: field // 搜索的字段
        });
        return false; // 阻止默认 form 跳转
    });

    exports('citys', {});
});