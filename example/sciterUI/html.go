package main

const html = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <style type="text/css">
        #form1 {
            text-align: center;
            height: 30px;
        }
        #p1,#p2 {
            font-size: 200%;
            white-space: pre
        }
        #tables,tr,td{
            font-max-size: 150%;
            width: 120px;
            border-collapse:collapse;
            border:1px solid #1F90FF;
        }
    </style>
    <title>奇门</title>
</head>

<body>

    <form id="form1" name="formymdh">
        <!-- 下拉年 -->
        <select id="yearid" name="year">
            <script type="text/tiscript">
                var t = new Date();
                var select = $(select#yearid);
                for( var i = 1601; i <= 3498; ++i ){
                    select.options.append(<option value={i}>{i}年</option>);
                    var s = select.text;
                    select.value=t.year;
                }
            </script>
        </select>
        <!-- 下拉月份 -->
        <select id="monthid" name="month">
            <script type="text/tiscript">
                var marr=["正月(寅)", "二月(卯)", "三月(辰)", "四月(巳)", "五月(午)", "六月(未)", "七月(申)", "八月(酉)", "九月(戌)", "十月(亥)", "十一月(子)", "十二月(丑)"]
                var selectm= $(select#monthid);
                for (var i =0;i<12;++i){
                    selectm.options.append(<option value={i+1}>{marr[i]}</option>);
                    var m = select.text;
                    selectm.value = t.month;
                }
            </script>
        </select>
        <!-- 下拉日期 -->
        <select id="dayid" name="day">
            <script type="text/tiscript">
                    var selectd= $(select#dayid);
                    var darr = ["初一", "初二", "初三", "初四", "初五", "初六", "初七", "初八", "初九", "初十",
                    "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九", "二十",
                    "廿一", "廿二", "廿三", "廿四", "廿五", "廿六", "廿七", "廿八", "廿九", "三十"];
                    for (var i =1;i<31;++i){
                        selectd.options.append(<option value={i}>{darr[i-1]}</option>);
                        var d = select.text;
                        if (t.day ==31){
                            selectd.value = t.day-1;
                        }else{
                            selectd.value = t.day;
                        }
                    }
            </script>
        </select>
        <!-- 时辰 子时1 丑时2 寅时3... -->
        <select id="hourid" name="hour">
            <script type="text/tiscript">
                    var selecth =$(select#hourid);
                    var harr=["子时(23~1)", "丑时(1~3)", "寅时(3~5)", "卯时(5~7)", "辰时(7~9)", "巳时(9~11)", "午时(11~13))", "未时(13~15)", "申时(15~17)", "酉时(17~19)", "戌时(19~21)","亥时(21~23)"]
                    for (var i =0;i<12;++i){
                      selecth.options.append(<option value={i+1}>{harr[i]}</option>);
                      var h = select.text;
                      if (t.hour/2==0){
                        selecth.value = 1;
                      }else{
                        selecth.value = t.hour/2;
                      }
                    }
            </script>
        </select>
        <!--生肖选择-->
        <select id="sxid" name="shengxiao">
            <script type="text/tiscript">
                    var selectsx =$(select#sxid);
                    var sxarr=["子水鼠","丑土牛","寅木虎","卯木兔","辰土龙","巳火蛇","午火马","未土羊","申金猴","酉金鸡","戌土狗","亥水猪"]
                    for (var i =0;i<12;++i){
                      selectsx.options.append(<option value={i+1}>{sxarr[i]}</option>);
                      var sx = select.text;
                      selectsx.value = 2;
                    }
            </script>
        </select>
        <!--闰月选择-->
        <select id="leapmbid" name="lmb">
            <script type="text/tiscript">
                    var selectmb=$(select#leapmbid);
                    var mbarr =["否","是"];
                    for (var i =0;i<2;++i){
                      selectmb.options.append(<option value={i}>{mbarr[i]}</option>);
                      var mb = select.text;
                      selectmb.value = 0;
                     }
            </script>
        </select>
        <button id="btn1">纪年信息</button>
        <button id="btn2" title="从上到下依次是九星->八门->八神->天盘奇仪->地盘奇仪->暗干支">奇门</button>
    </form>

    <p id="p1"></p>
    <p id="p2"></p>

    <table id="tables" rules="all" width="100%">
            <tr id="492">
                <td id="4" style="font-size: 20px"></td>
                <td id="9"style="font-size: 20px"></td>
                <td id="2"style="font-size: 20px"></td>
            </tr>
            <tr id="357">
                <td id="3"style="font-size: 20px"></td>
                <td id="5"style="font-size: 20px"></td>
                <td id="7"style="font-size: 20px"></td>
            </tr>
            <tr id="816">
                <td id="8"style="font-size: 20px"></td>
                <td id="1"style="font-size: 20px"></td>
                <td id="6"style="font-size: 20px"></td>
            </tr>
    </table>
    <script type="text/tiscript">
        function ymdhsb(){
              //年
              var opt = $(select[name='year']).$$(option);
              var ly= 0;
              for(var child in opt) {
                  if(child.getState(Element.STATE_CHECKED)) {
                      ly = child.value;
                  }
              }
              //月
              var optm = $(select[name='month']).$$(option);
              var lm= 0;
              for(var child in optm) {
                  if(child.getState(Element.STATE_CHECKED)) {
                      lm = child.value;
                  }
              }
              //日
              var optd = $(select[name='day']).$$(option);
              var ld= 0;
              for(var child in optd) {
                  if(child.getState(Element.STATE_CHECKED)) {
                      ld = child.value;
                  }
              }
               //时辰
              var opth = $(select[name='hour']).$$(option);
              var lh= 0;
              for(var child in opth) {
                  //判断元素是否选中
                  if(child.getState(Element.STATE_CHECKED)) {
                      lh = child.value;
                    //  stdout.println("选中h:",lh);
                  }
              }
              //生肖
              var optsx = $(select[name='shengxiao']).$$(option);
              var sx =0;
              for(var child in optsx){
              //判断元素是否选中
                  if(child.getState(Element.STATE_CHECKED)){
                      sx = child.value;
                     // stdout.println("选中sx:",sx);
                  }
              }
              //闰月
              var optmb=$(select[name='lmb']).$$(option);
              var lb=0;
              for(var child in optmb){
                  if(child.getState(Element.STATE_CHECKED)){
                      lb = child.value;
                  }
              }
             var formData = {
             "year":ly,
             "month":lm,
             "day":ld,
             "hour":lh,
             "zodiac":sx,
             "leapmb":lb
             };
             return formData;
        }
        $(#btn1).on("click",function(){
            var dates = ymdhsb();
            var slgn = view.ymdinfo(dates.year,dates.month, dates.day, dates.hour,dates.zodiac, dates.leapmb);
            var jsData = parseData(slgn);
            if(jsData){
                $(#p1).html = jsData.solar+"<br>"+jsData.lunar+"<br>"+jsData.gan_zhi+"<br>"+jsData.na_yin+"<br>";
            }else{
                stdout.println("no data");
            }
        });
        $(#btn2).on("click",function(){
              var infoymdh = ymdhsb();
              var info = view.qimeninfo(infoymdh.year,infoymdh.month, infoymdh.day, infoymdh.hour,infoymdh.zodiac, infoymdh.leapmb);
              var qmjs = parseData(info);
              if(qmjs){
                $(#p2).html =qmjs.jie_qi+" "+qmjs.yin_yang+" "+qmjs.n+"局"+" "+qmjs.yuan+"<br>"+"值符:"+qmjs.zhifu+" "+"值使:"+qmjs.zhishi+"<br>";
                $(#4).html = qmjs.g_4[0]+"<br>"+ qmjs.g_4[1]+"<br>"+ qmjs.g_4[4]+"<br>"+ qmjs.g_4[3]+"<br>"+ qmjs.g_4[5]+"<br>"+ qmjs.g_4[2]+"<br>";
                $(#9).html = qmjs.g_9[0]+"<br>"+ qmjs.g_9[1]+"<br>"+ qmjs.g_9[4]+"<br>"+ qmjs.g_9[3]+"<br>"+ qmjs.g_9[5]+"<br>"+ qmjs.g_9[2]+"<br>";
                $(#2).html = qmjs.g_2[0]+"<br>"+ qmjs.g_2[1]+"<br>"+ qmjs.g_2[4]+"<br>"+ qmjs.g_2[3]+"<br>"+ qmjs.g_2[5]+"<br>"+ qmjs.g_2[2]+"<br>";
                $(#3).html =qmjs.g_3[0]+"<br>"+ qmjs.g_3[1]+"<br>"+ qmjs.g_3[4]+"<br>"+ qmjs.g_3[3]+"<br>"+ qmjs.g_3[5]+"<br>"+ qmjs.g_3[2]+"<br>";
                $(#5).html = qmjs.g_5[0]+"<br>"+ qmjs.g_5[1]+"<br>"+ qmjs.g_5[4]+"<br>"+ qmjs.g_5[3]+"<br>"+ qmjs.g_5[5]+"<br>"+ qmjs.g_5[2]+"<br>";
                $(#7).html = qmjs.g_7[0]+"<br>"+ qmjs.g_7[1]+"<br>"+ qmjs.g_7[4]+"<br>"+ qmjs.g_7[3]+"<br>"+ qmjs.g_7[5]+"<br>"+ qmjs.g_7[2]+"<br>";
                $(#1).html = qmjs.g_1[0]+"<br>"+ qmjs.g_1[1]+"<br>"+ qmjs.g_1[4]+"<br>"+ qmjs.g_1[3]+"<br>"+ qmjs.g_1[5]+"<br>"+ qmjs.g_1[2]+"<br>";
                $(#8).html = qmjs.g_8[0]+"<br>"+ qmjs.g_8[1]+"<br>"+ qmjs.g_8[4]+"<br>"+ qmjs.g_8[3]+"<br>"+ qmjs.g_8[5]+"<br>"+ qmjs.g_8[2]+"<br>";
                $(#6).html = qmjs.g_6[0]+"<br>"+ qmjs.g_6[1]+"<br>"+ qmjs.g_6[4]+"<br>"+ qmjs.g_6[3]+"<br>"+ qmjs.g_6[5]+"<br>"+ qmjs.g_6[2]+"<br>";
              }else{
                stdout.println("No data");
              }
        });
    </script>
</body>
</html>
`
