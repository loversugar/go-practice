// pages/chat/chat.js
var inputVal = '';
var keyHeight = 0;
var sockTask = {};
Page({

  /**
   * 页面的初始数据
   */
  data: {
    inputBottom: 0,
    inputVal: '',
    messageArray: [{"name": "message", "password": "hah"}]
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this._changeNavigationBarTitle(options.name); 
    // this._connectSocket();
    // sockTask.onMessage(result => {
    //   console.log(result) 
    // });
  },

  openConnect() {
    this._connectSocket();
    sockTask.onMessage(result => {
      this.setData({
        messageArray: this.data.messageArray.concat(JSON.parse(result.data))
      })
      console.log(this.data.messageArray);
    });
  },

  _changeNavigationBarTitle(name) {
    wx.setNavigationBarTitle({
      title: name
    })
  },

  _connectSocket() {
    sockTask = wx.connectSocket({
      url: 'ws://localhost:8080/echo'
    });
    sockTask.onOpen(result => {
      console.log(result); 
    });
  },

   /**
   * 获取聚焦
   */
  focus: function(event) {
    this.setData({
      inputBottom: keyHeight
    })
  },

  //失去聚焦(软键盘消失)
  blur: function(e) {
    this.setData({
      inputBottom: 0
    })
  },

  sendMessage: function() {
    sockTask.send({
      data: JSON.stringify({name: "zs", password: "pwd"}),
      success: (result)=>{
      },
      fail: ()=>{},
      complete: ()=>{}
    });
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})