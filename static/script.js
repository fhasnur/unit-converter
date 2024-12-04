const openConverter = (evt, converter) => {
  const tabcontent = document.getElementsByClassName("tabcontent");
  const tablinks = document.getElementsByClassName("tablinks");

  for (let i = 0; i < tabcontent.length; i++) {
    tabcontent[i].style.display = "none";
  }

  for (let i = 0; i < tablinks.length; i++) {
    tablinks[i].className = tablinks[i].className.replace(" active", "");
  }

  const selectedTab = document.getElementById(converter);
  if (selectedTab) {
    selectedTab.style.display = "block";
    evt.currentTarget.className += " active";
  } else {
    console.error(`Tab with ID "${converter}" not found.`);
  }
};

const resetResult = () => {
  const contentResults = document.querySelectorAll(".result");
  contentResults.forEach((result) => (result.style.display = "none"));
};

document.addEventListener("DOMContentLoaded", () => {
  document.getElementById("defaultOpen")?.click();
});