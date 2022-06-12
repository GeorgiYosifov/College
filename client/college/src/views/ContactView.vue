<template>
  <section id="" class="contact">
    <h1 class="title">Contact with us</h1>
    <article class="contact-section">
      <article class="contacts">
        <div class="contacts-card">
          <div class="icon">
            <i class="fa-solid fa-location-dot"></i>
          </div>
          <div class="text">
            <h3 class="text-title">Address</h3>
            <p>Sofia, quarter Ivan Vazov ,1</p>
          </div>
        </div>
        <div class="contacts-card">
          <div class="icon">
            <i class="fa-solid fa-phone"></i>
          </div>
          <div class="text">
            <h3 class="text-title">Phone</h3>
            <p><a href="tel://1234567920">02 / 789 87 65</a></p>
          </div>
        </div>
        <div class="contacts-card">
          <div class="icon">
            <i class="fa-solid fa-envelope"></i>
          </div>
          <div class="text">
            <h3 class="text-title">Email</h3>
            <p><a href="mailto:techcollege@gmail.com">techcollege@gmail.com</a></p>
          </div>
        </div>
      </article>
      <article class="contact-section-content">
        <article class="contact-form-contect">
          <h3 class="contact-form-contect-title">Ask your question</h3>
          <form method="" id="" class="contact-form">
            <div class="input-firse-name">
              <label for="fname">First Name</label>
              <input type="text" id="fname" name="firstname" placeholder="Your name.." required>
            </div>
            <div class="input-last-name">
              <label for="lname">Last Name</label>
              <input type="text" id="lname" name="lastname" placeholder="Your last name.." required>
            </div>

            <label for="subject">Subject</label>
            <textarea id="subject" name="subject" placeholder="Write something.."
                      style="height:200px"></textarea>

            <div class="contact-button">
              <button class="btn">Sent</button>
            </div>
          </form>
        </article>
        <article class="contact-section-image">
          <img src="@/assets/img/contact.png" />
        </article>
      </article>
    </article>
  </section>
</template>

