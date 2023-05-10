const https = require('https')

https.get('https://push2.eastmoney.com/api/qt/ulist.np/get?fltt=2&fields=f12,f13,f19,f14,f139,f148,f2,f4,f1,f125,f18,f3,f152,f5,f30,f31,f32,f6,f8,f7,f10,f22,f9,f112,f100,f88,f153&secids=1.000001,0.399001,0.399006,1.000688,100.HSI', res => {
    let list = [];
    res.on('data', chunk => {
        list.push(chunk);
    });
    res.on('end', () => {
        const { data } = JSON.parse(Buffer.concat(list).toString());
      console.log('data>>>', data);
    });
}).on('error', err => {
    console.log('Error: ', err.message);
});