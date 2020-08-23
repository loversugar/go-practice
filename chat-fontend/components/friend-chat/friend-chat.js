Component({
  data: {
    isTouchChange: false
  },
  properties: {
    chatInfo: Object
  },
  methods: {
    onNavigateToChat() {
      wx.navigateTo({
        url: `/pages/chat/chat?friendId=${this.data.chatInfo.id}&name=${this.data.chatInfo.name}`,
        success: (result)=>{
          
        },
        fail: ()=>{},
        complete: ()=>{}
      });
    },

    onTouchStart() {
      this.setData({
        isTouchChange: true
      })
    },

    onTouchEnd() {
      this.setData({
        isTouchChange: false
      })
    }
  }
})