// pages/chat/chat.js
var inputVal = '';
var keyHeight = 0;
Page({

  /**
   * 页面的初始数据
   */
  data: {
    inputBottom: 0
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    this._changeNavigationBarTitle(options.name); 
    this._connectSocket();
  },

  _changeNavigationBarTitle(name) {
    wx.setNavigationBarTitle({
      title: name
    })
  },

  _connectSocket() {
    var sockTask = wx.connectSocket({
      url: 'ws://localhost:8080/echo'
    });
    sockTask.onOpen(result => {
      console.log(result); 

      sockTask.send({
        data: JSON.stringify({"name": "zs", "password": "pwd"}),
        success: (result)=>{
        },
        fail: ()=>{},
        complete: ()=>{}
      });
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