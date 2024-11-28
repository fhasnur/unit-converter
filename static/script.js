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

function resetForm() {
  document.querySelector('.result').style.display = 'none';
  document.querySelector('.btn-reset').style.display = 'none';
}

document.getElementById("defaultOpen")?.click();