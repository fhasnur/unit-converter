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

    localStorage.setItem("activeTab", converter);
  } else {
    console.error(`Tab with ID "${converter}" not found.`);
  }
};

const resetResult = () => {
  const contentResults = document.querySelectorAll(".result");
  contentResults.forEach((result) => (result.style.display = "none"));
};

document.addEventListener("DOMContentLoaded", () => {
  const activeTab = localStorage.getItem("activeTab") || 'length';

  if (activeTab) {
    const activeTabElement = document.getElementById(activeTab);
    const activeTabButton = document.querySelector(`button[onclick*="${activeTab}"]`);

    if (activeTabElement && activeTabButton) {
      activeTabElement.style.display = "block";
      activeTabButton.classList.add("active");
    }
  }
});