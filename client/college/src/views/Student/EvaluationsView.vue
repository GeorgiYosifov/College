<section id="" class="login-students">
  <article class="enrollment-semestar">
    <h1 class="secondary-title">Evaluations</h1>
    <div class="semestar">
      <h4 class="semestar">Semester:</h4>
      <div class="evaluations-select">
        <select v-model="selected">
          <option value="">Choose semester:</option>
          <option value="1">First year, I semester</option>
          <option value="2">First year, II semester</option>
          <option value="3">Second year, I semester</option>
        </select>
      </div>
    </div>
    <table id="" class="table-informations">
      <thead>
        <tr>
          <th>Course code</th>
          <th>Course</th>
          <th>Evaluations</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>001</td>
          <td>Linear algebra</td>
          <td>T1 (Test work) <br>
            T2 (Test work) <br>
            T3 (Presence and participation) <br>
            TO (Final assessment) <br>
            I1 (Exam)
          </td>
        </tr>
      </tbody>
    </table>
  </article>
</section>