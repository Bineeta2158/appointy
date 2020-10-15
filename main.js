$(document).ready(function(){
  $('.swipe').slick({
    autoplay: false,
    arrows: false,
    prevArrow: $('.prev'),
    nextArrow: $('.next'),
    dots: false,
      pauseOnHover: false,
      responsive: [{
      breakpoint: 768,
      settings: {
        slidesToShow: 1
      }
    }, {
      breakpoint: 520,
      settings: {
        slidesToShow: 1
      }
    }]
  });
});
