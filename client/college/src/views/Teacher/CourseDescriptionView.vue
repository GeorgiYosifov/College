<template>
    <nav class="nav" id="idnav">
        <ul class="secondary-menu">
            <li class="nav-links">
                <RouterLink to="/teacherEvaluationsCourse"><span>Evaluations</span></RouterLink>
            </li>
            <li class="nav-links">
                <RouterLink to="/teacherCourseDescription"><span>Course description</span></RouterLink>
            </li>
        </ul>
    </nav>

    <section id="" class="course-section">
        <article class="course-content">
            <h3 class="course-content-title">Linear algebra</h3>
            <p class="course-description-paragraph">
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry.
                Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer
                took
                a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries,
                but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in
                the
                1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with
                desktop
                publishing software like Aldus PageMaker including versions of Lorem Ipsum.
                <i class="fa-solid fa-pen edit"></i>
            </p>
            <p class="course-description-paragraph"><span>Teachers:</span> Tom Pit <i class="fa-solid fa-pen edit"></i></p>
            <br>
            <p class="course-description-paragraph">   <span>Acquire knowledge: </span>    <i class="fa-solid fa-pen edit"></i> </p>
              <ul class="course-evaluation-list">
              <li>t1</li>
              <li>t2</li>
              <li>t3</li>
            </ul>
            <p class="course-description-paragraph"> <span> Prerequisites: </span> <i class="fa-solid fa-pen edit"></i></p>
            <ul class="course-evaluation-list">
              <li>t1</li>
              <li>t2</li>
              <li>t3</li>
            </ul>
            <p class="course-description-paragraph">  <span>Topics for consideration in the course: </span> <i class="fa-solid fa-pen edit"></i>  </p>
            <ol class="course-evaluation-list">
              <li>t1</li>
              <li>t2</li>
              <li>t3</li>
            </ol>
            <p class="course-description-paragraph"><span>Evaluation: </span><i class="fa-solid fa-pen edit"></i></p>
            <ul class="course-evaluation-list">
              <li>t1</li>
              <li>t2</li>
              <li>t3</li>
            </ul>
            <button class="btn">Save</button>
        </article>
    </section>
</template>
